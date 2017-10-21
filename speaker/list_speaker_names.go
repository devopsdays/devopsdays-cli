package speaker

func listSpeakerNames(speakers []string, city string, year string) (speakerFullNames []string, err error) {
	for _, f := range speakers {
		var mySpeaker Speaker
		mySpeaker, err = GetSpeakerInfo(f, city, year)
		speakerFullNames = append(speakerFullNames, mySpeaker.Title)
	}
	return
}
