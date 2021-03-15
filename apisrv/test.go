package main

import (
	"fmt"
	xgin "github.com/gin-gonic/gin"
	ego2 "github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
	"github.com/jcuga/golongpoll"
	"math/rand"
	longpoll2 "roomx/components/longpoll"
	"time"
)

func main() {
	ego := ego2.New()
	gin := egin.Load("server.http").Build()
	longpoll := longpoll2.Load("server.longpoll").Build(gin)

	go generateRandomEvents(longpoll.LongpollManager())

	gin.GET("/basic/events", func(ctx *xgin.Context) {
		longpoll.LongpollManager().SubscriptionHandler(ctx.Writer, ctx.Request)
	})
	gin.GET("/basic", BasicExampleHomepage)
	if err := ego.Serve(longpoll).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func generateRandomEvents(lpManager *golongpoll.LongpollManager) {
	farmEvents := []string{
		"Cow says 'Moooo!'",
		"Duck went 'Quack!'",
		"Chicken says: 'Cluck!'",
		"Goat chewed grass.",
		"Pig went 'Oink! Oink!'",
		"Horse ate hay.",
		"Tractor went: Vroom Vroom!",
		"Farmer ate bacon.",
	}
	// every 0-5 seconds, something happens at the farm:
	for {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		lpManager.Publish("farm", farmEvents[rand.Intn(len(farmEvents))])
	}
}

func BasicExampleHomepage(ctx *xgin.Context) {
	fmt.Fprintf(ctx.Writer, `
<html>
<head>
    <title>golongpoll basic example</title>
</head>
<body>
    <h1>golongpoll basic example</h1>
    <h2>Here's whats happening around the farm:</h2>
    <ul id="animal-events"></ul>
<script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
<script>

    // for browsers that don't have console
    if(typeof window.console == 'undefined') { window.console = {log: function (msg) {} }; }

    // Start checking for any events that occurred after page load time (right now)
    // Notice how we use .getTime() to have num milliseconds since epoch in UTC
    // This is the time format the longpoll server uses.
    var sinceTime = (new Date(Date.now())).getTime();

    // Let's subscribe to animal related events.
    var category = "farm";

    (function poll() {
        var timeout = 45;  // in seconds
        var optionalSince = "";
        if (sinceTime) {
            optionalSince = "&since_time=" + sinceTime;
        }
        var pollUrl = "/basic/events?timeout=" + timeout + "&category=" + category + optionalSince;
        // how long to wait before starting next longpoll request in each case:
        var successDelay = 10;  // 10 ms
        var errorDelay = 3000;  // 3 sec
        $.ajax({ url: pollUrl,
            success: function(data) {
                if (data && data.events && data.events.length > 0) {
                    // got events, process them
                    // NOTE: these events are in chronological order (oldest first)
                    for (var i = 0; i < data.events.length; i++) {
                        // Display event
                        var event = data.events[i];
                        $("#animal-events").append("<li>" + event.data + " at " + (new Date(event.timestamp).toLocaleTimeString()) +  "</li>")
                        // Update sinceTime to only request events that occurred after this one.
                        sinceTime = event.timestamp;
                    }
                    // success!  start next longpoll
                    setTimeout(poll, successDelay);
                    return;
                }
                if (data && data.timeout) {
                    console.log("No events, checking again.");
                    // no events within timeout window, start another longpoll:
                    setTimeout(poll, successDelay);
                    return;
                }
                if (data && data.error) {
                    console.log("Error response: " + data.error);
                    console.log("Trying again shortly...")
                    setTimeout(poll, errorDelay);
                    return;
                }
                // We should have gotten one of the above 3 cases:
                // either nonempty event data, a timeout, or an error.
                console.log("Didn't get expected event data, try again shortly...");
                setTimeout(poll, errorDelay);
            }, dataType: "json",
        error: function (data) {
            console.log("Error in ajax request--trying again shortly...");
            setTimeout(poll, errorDelay);  // 3s
        }
        });
    })();
</script>
</body>
</html>`)
}