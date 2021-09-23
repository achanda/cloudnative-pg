/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2021 EnterpriseDB Corporation.
*/

package tests

// List of the labels we use for labeling test specs
// See https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md#label-decoration
const (
	// LabelDisruptive is the string for labelling disruptive tests
	LabelDisruptive = "disruptive"

	// LabelPerformance is the string for labelling performance tests
	LabelPerformance = "performance"

	// LabelUpgrade is the string for labelling upgrade tests
	LabelUpgrade = "upgrade"
)