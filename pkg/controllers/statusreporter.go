package controllers

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/util/version"

	"github.com/operator-framework/operator-sdk/pkg/status"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"

	aro "github.com/Azure/ARO-RP/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
)

type StatusReporter struct {
	arocli aroclient.AroV1alpha1Interface
	name   types.NamespacedName
	log    *logrus.Entry
}

var (
	emptyConditions = []status.Condition{
		{
			Type:    aro.InternetReachable,
			Status:  corev1.ConditionUnknown,
			Reason:  "",
			Message: "",
		},
	}
)

func NewStatusReporter(log *logrus.Entry, arocli aroclient.AroV1alpha1Interface, namespace, name string) *StatusReporter {
	return &StatusReporter{
		log:    log.WithField("manager", "StatusReporter"),
		arocli: arocli,
		name:   types.NamespacedName{Name: name, Namespace: namespace},
	}
}

func historyShorten(current []aro.ReconsileHistory, limit int) []aro.ReconsileHistory {
	histLen := len(current)
	if histLen < 11 {
		return current
	}
	return current[histLen-10 : histLen]
}

func (r *StatusReporter) AddReconcileAction(ctx context.Context, action, reason string) error {
	time := metav1.Now()

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		co, err := r.arocli.Clusters(r.name.Namespace).Get(r.name.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		co.Status.ReconsileHistory = historyShorten(append(co.Status.ReconsileHistory,
			aro.ReconsileHistory{
				CompletionTime: &time,
				Action:         action,
				Reason:         reason}), 10)

		setStaticStatus(&co.Status, r.name.Namespace)

		_, err = r.arocli.Clusters(r.name.Namespace).Update(co)
		return err
	})
}

func (r *StatusReporter) SetNoInternetConnection(ctx context.Context, connectionErr error) error {
	time := metav1.Now()
	msg := "Outgoing connection failed"
	if connectionErr != nil {
		msg += ": " + connectionErr.Error()
	}
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		co, err := r.arocli.Clusters(r.name.Namespace).Get(r.name.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		co.Status.Conditions.SetCondition(status.Condition{
			Type:               aro.InternetReachable,
			Status:             corev1.ConditionFalse,
			Message:            msg,
			Reason:             "CheckFailed",
			LastTransitionTime: time})

		setStaticStatus(&co.Status, r.name.Namespace)

		_, err = r.arocli.Clusters(r.name.Namespace).Update(co)
		return err
	})
}

func (r *StatusReporter) SetInternetConnected(ctx context.Context) error {
	time := metav1.Now()
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		co, err := r.arocli.Clusters(r.name.Namespace).Get(r.name.Name, v1.GetOptions{})
		if err != nil {
			return err
		}

		co.Status.Conditions.SetCondition(status.Condition{
			Type:               aro.InternetReachable,
			Status:             corev1.ConditionTrue,
			Message:            "Outgoing connection successful.",
			Reason:             "CheckDone",
			LastTransitionTime: time})

		setStaticStatus(&co.Status, r.name.Namespace)

		_, err = r.arocli.Clusters(r.name.Namespace).Update(co)
		return err
	})
}

func setStaticStatus(status *aro.ClusterStatus, namespace string) {
	if len(status.RelatedObjects) == 0 {
		status.RelatedObjects = []corev1.ObjectReference{
			{Kind: "Namespace", Name: namespace},
			{Kind: "Secret", Name: "pull-secret", Namespace: "openshift-config"},
		}
	}
	status.OperatorVersion = version.GitCommit
}
