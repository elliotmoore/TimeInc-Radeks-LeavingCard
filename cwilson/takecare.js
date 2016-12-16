


var http = require('http');

const PORT=3000; 

const since = new Date('2016-12-16T17:00:00.123Z');

function handleRequest(request, response){
    response.end('It\'s been '+daysSince()+' days since you left... ' +
    	'and that completly sucks!!! \nIt was a pleasure working with you wishing you only ' +
    	'good fourtune for your future. \nBest of luck Radek! Keep up the good fight... IaC!');
}

var server = http.createServer(handleRequest);

server.listen(PORT, function(){
    console.log("curl http://localhost:%s", PORT);
});

function daysSince() {
    var ONE_DAY = 1000 * 60 * 60 * 24
    var date1_ms = since.getTime()
    var date2_ms = new Date()
    var difference_ms = Math.abs(date1_ms - date2_ms)
    return Math.round(difference_ms/ONE_DAY)

}