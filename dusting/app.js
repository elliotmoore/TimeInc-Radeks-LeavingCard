var prompt = require('prompt');
var sleep = require('sleep');

  var properties = [
    {
      name: 'Did you get better pay?',
      validator: /yes/,
      warning: 'You cannot leave unless it is better'
    },
    {
      name: 'Is it a better job?',
      validator: /yes/,
      warning: 'You cannot leave unless it is better'
    },
    {
      name: 'Are you leaving good friends?'
    },
  ];

  prompt.start();

  prompt.get(properties, function (err, result) {
    if (err) { return onErr(err); }
    if(result['Are you leaving good friends?'].toUpperCase() === 'YES'){
      console.log('Good Luck you will be missed!')
    }else{
      var text = 'Good riddence!';
      console.log('Here is a script loop for your enjoyment');
      sleep.sleep(2);
      console.log(message(text));
    }
  });

  function onErr(err) {
    console.log(err);
    return 1;
  }

  function message(text){
    console.log(text);
    sleep.sleep(1);
    message(text);
  }
