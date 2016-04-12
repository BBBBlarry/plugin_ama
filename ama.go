package ama



import (
    "log"
    "net/http"
    "io/ioutil"
    "github.com/itsabot/abot/shared/datatypes"
    "github.com/itsabot/abot/shared/nlp"
    "github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
    // Abot should route messages to this plugin that contain any combination
    // of the below words. The stems of the words below are used, so you don't
    // need to include duplicates (e.g. there's no need to include both "stock"
    // and "stocks"). Everything will be lowercased as well, so there's no
    // difference between "ETF" and "etf".
    trigger := &nlp.StructuredInput{
        Commands: []string{"yo", "hey"},
        Objects: []string{"wiki"},
    }


    // Tell Abot how this plugin will respond to new conversations and follow-up
    // requests.
    fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}

    // Create the plugin.
    var err error
    pluginPath := "github.com/BBBBlarry/plugin_ama"
    p, err = plugin.New(pluginPath, trigger, fns)
    if err != nil {
        log.Fatalln("building", err)
    }

    p.Vocab = dt.NewVocab(
        dt.VocabHandler{
            Fn: kwQueryAlpha,
            Trigger: &nlp.StructuredInput{
                Commands: []string{"yo", "hey"},
                Objects: []string{"wiki"},
            },
        },
    )
}
    

// Abot calls Run the first time a user interacts with a plugin
func Run(in *dt.Msg) (string, error) {
    return FollowUp(in)
}

// Abot calls FollowUp every subsequent time a user interacts with the plugin
// as long as the messages hit this plugin consecutively. As soon as Abot sends
// a message for this user to a different plugin, this plugin's Run function
// will be called the next it's triggered.  This Run/FollowUp design allows us
// to reset a plugin's state when a user changes conversations.
func FollowUp(in *dt.Msg) (string, error) {
    //return QueryAlpha(in), nil
   return p.Vocab.HandleKeywords(in), nil
}


func kwQueryAlpha(in *dt.Msg) (resp string){
    res, err := http.Get("http://www.google.com/robots.txt")
    if err != nil {
        log.Fatal(err)
    }
    robots, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    return string(robots)

}

