package migration

func upgradeDatabaseToVersion3_1_0(fromVersion string) error {
	if UpdateErr := updateSchemaVersion(version3_1_0); UpdateErr != nil {
		return UpdateErr
	}
	return nil
}

func upgradeDatabaseToVersion3_1_1(fromVersion string) error {
	if UpdateErr := updateSchemaVersion(version3_1_1); UpdateErr != nil {
		return UpdateErr
	}
	return nil
}
