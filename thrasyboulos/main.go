package main

import (
	"github.com/kpango/glg"
	"github.com/odysseia-greek/knossos/thrasyboulos/app"
	"github.com/odysseia-greek/knossos/thrasyboulos/config"
	"os"
	"strings"
	"time"
)

func init() {
	errlog := glg.FileWriter("/tmp/error.log", 0666)
	defer errlog.Close()

	glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.ERR, errlog)
}

func main() {
	//https://patorjk.com/software/taag/#p=display&f=Crawford2&t=THRASYBOULOS
	glg.Info("\n ______  __ __  ____    ____  _____ __ __  ____    ___   __ __  _       ___   _____\n|      ||  |  ||    \\  /    |/ ___/|  |  ||    \\  /   \\ |  |  || |     /   \\ / ___/\n|      ||  |  ||  D  )|  o  (   \\_ |  |  ||  o  )|     ||  |  || |    |     (   \\_ \n|_|  |_||  _  ||    / |     |\\__  ||  ~  ||     ||  O  ||  |  || |___ |  O  |\\__  |\n  |  |  |  |  ||    \\ |  _  |/  \\ ||___, ||  O  ||     ||  :  ||     ||     |/  \\ |\n  |  |  |  |  ||  .  \\|  |  |\\    ||     ||     ||     ||     ||     ||     |\\    |\n  |__|  |__|__||__|\\_||__|__| \\___||____/ |_____| \\___/  \\__,_||_____| \\___/  \\___|\n                                                                                   \n")
	glg.Info(strings.Repeat("~", 37))
	glg.Info("\"πέμψας γὰρ παρὰ Θρασύβουλον κήρυκα ἐπυνθάνετο ὅντινα ἂν τρόπον ἀσφαλέστατον καταστησάμενος τῶν πρηγμάτων κάλλιστα τὴν πόλιν ἐπιτροπεύοι.\"")
	glg.Info("\"He had sent a herald to Thrasybulus and inquired in what way he would best and most safely govern his city. \"")
	glg.Info(strings.Repeat("~", 37))

	glg.Debug("creating config")

	env := os.Getenv("ENV")

	thrasyboulosConfig, err := config.CreateNewConfig(env)
	if err != nil {
		glg.Error(err)
		glg.Fatal("death has found me")
	}

	duration := time.Millisecond * 5000
	minute := time.Minute * 60
	timeFinished := minute.Milliseconds()

	done := make(chan bool)
	handler := app.ThrasyboulosHandler{Config: thrasyboulosConfig, Duration: duration, TimeFinished: timeFinished}

	go func() {
		handler.WaitForJobsToFinish(done)
	}()

	select {

	case <-done:
		glg.Infof("%s job finished", thrasyboulosConfig.Job)
		os.Exit(0)
	}

}
