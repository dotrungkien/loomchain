package gen

/*
var (
	removeTestDirectory = true
	ip = "127.0.0.1:10000"
)


func add(mockService *httpmock.MockHTTPServer, spinUrl string, testFile string) error {
	// define request->response pairs
	requestUrl, _ := url.Parse(spinUrl)
	raw, err := ioutil.ReadFile(testFile)
	if err != nil {
		return err
	}
	mockService.AddResponses([]httpmock.MockResponse{
		{
			Request: http.Request{
				Method: "GET",
				URL:    requestUrl,
			},
			Response: httpmock.Response{
				StatusCode: 200,
				Body:       string(raw),
			},
		},
	})
	return nil
}

func TestSpin(t *testing.T) {
	type spinTestParms struct {
		spinUrl, outDir, name, dataFile string
	}
	testDir, err := ioutil.TempDir("", "testspin")
	if err !=  nil {
		t.Errorf("error creating test directory")
	}
	if removeTestDirectory {
		defer os.RemoveAll(testDir)
	}
	fmt.Println("Using test directory ", testDir)

	spins := []spinTestParms{
		{
			"http://127.0.0.1:10000/github.com/loomnetwork/testproj1-core/archive/master.zip",
			testDir, "",
			"./testdata/testproj-core-master.zip",
		},
		{
			"http://127.0.0.1:10000/github.com/loomnetwork/testproj2-core/archive/master.zip",
			testDir, "mytestproject",
			"./testdata/testproj-core-master.zip",
		},
		{
			"http://127.0.0.1:10000/github.com/loomnetwork/weave-testproj3-core/archive/master.zip",
			testDir, "",
			"./testdata/weave-testproj-core-master.zip",
		},
		{
			"http://127.0.0.1:10000/github.com/loomnetwork/weave-testproj4-core/archive/master.zip",
			testDir, "anothertestproj",
			"./testdata/weave-testproj-core-master.zip",
		},
	}

	mockService := httpmock.NewMockHTTPServer(ip)

	for _, tests := range spins {

		add(mockService, tests.spinUrl, tests.dataFile)

		spinTitle, _, err := getRepoPath(tests.spinUrl)
		if err != nil {
			t.Error("bad repoPath")
		}
		projName := projectName(tests.name, spinTitle)
		willCreateDir := filepath.Join(getOutDir(tests.outDir), projName)

		err = Spin(tests.spinUrl, tests.outDir, tests.name)
		if err != nil {
			fmt.Println(err)
			t.Error("something went wrong with spinning %s, %s, %s", tests.spinUrl, tests.outDir, tests.name)
		}
		if _, err := os.Stat(willCreateDir); err != nil {
			t.Error("has not made directory %s", willCreateDir)
		}
	}

}
*/
