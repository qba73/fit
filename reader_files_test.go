// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		false,
		11839585052621994073,
		true,
		tdoAllWithNullLogger,
	},
	{
		"fitsdk",
		"Activity.fit",
		false,
		11751796383273411430,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		false,
		8432119481642341932,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"Settings.fit",
		false,
		15312295986705592262,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		false,
		15637902928121490784,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		false,
		10030359812057747231,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		false,
		14879051105853803529,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		8356463518492741917,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		false,
		3776887550528028307,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		false,
		8168113212172708718,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		false,
		15838312440114176145,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		false,
		0,
		false,
		tdoNone,
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		false,
		8067952813397841073,
		true,
		tdoNone,
	},
	{
		"sram",
		"Settings.fit",
		false,
		7436219522300668883,
		true,
		tdoNone,
	},
	{
		"sram",
		"Settings2.fit",
		false,
		16468149498896926724,
		true,
		tdoNone,
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		4621767084268372354,
		true,
		tdoNone,
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		false,
		5100394846405776062,
		true,
		tdoNone,
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		false,
		11746850638063494813,
		true,
		tdoNone,
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		true,
		14351461695201116882,
		true,
		tdoNone,
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		true,
		10272597984625945189,
		true,
		tdoNone,
	},
}
