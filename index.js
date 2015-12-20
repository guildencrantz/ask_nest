// https://gist.github.com/miksago/d1c456d4e235e025791d

var child_process = require('child_process');

exports.handler = function(event, context) {
  var proc = child_process.spawn('./ask_nest', [ JSON.stringify(event) ]);

	var stdout = '';
	proc.stdout.on('data', function(buf) {
		stdout += buf;
	});

  proc.on('close', function(code){
    if(code !== 0) {
      return context.done(new Error("Process exited with non-zero status code"));
    }

    context.succeed(JSON.parse(stdout));
  });
}
