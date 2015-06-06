  var socket;
  var url_websocket = "127.0.0.1:8080";
  var url_tutorial;
  var url_code;

  var tutname='run';
  var prev='';
  var next='';

  $(document).ready(function () {
      //============ Create a socket
      //socket = new WebSocket('ws://' + window.location.host + '/console/sync?tutname=' + $('#tutname').text());
      socket = new WebSocket("ws://" + url_websocket + "/v1/testbuild/");

      $("#nav-tutorials").click();
      $("#nav-debug").addClass("active")
      //$(".dropdown-menu li a")[2].click();


  var terminal = new Terminal('terminal-panel', {theme: 'interlaced'}, {
    execute: function(cmd, args) {
        var fullcommand = cmd;
        for(var idx in args){
            fullcommand += " " + args[idx];
        }
        socket.send(fullcommand);
        switch (cmd) {
            case 'clear':
                terminal.clear();
                return '';

            case 'help':
                return 'Commands: clear, help, theme, ver or version<br>More help available <a class="external" href="http://github.com/SDA/terminal" target="_blank">here</a>';

            case 'theme':
                if (args && args[0]) {
                    if (args.length > 1) return 'Too many arguments';
                    else if (args[0].match(/^interlaced|modern|white$/)) { terminal.setTheme(args[0]); return ''; }
                    else return 'Invalid theme';
                }
                return terminal.getTheme();

            case 'ver':
            case 'version':
                return '1.0.0';

            default:
                // Unknown command.
                return "";
        };
    }
});

  //============ code mirror editor 
  var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    lineNumbers: true,
    theme: "lesser-dark",
    scrollbarStyle: "overlay",
    extraKeys: {
      "F11": function(cm) {
        cm.setOption("fullScreen", !cm.getOption("fullScreen"));
      },
      "Esc": function(cm) {
        if (cm.getOption("fullScreen")) cm.setOption("fullScreen", false);
      }
    }
  });

  $('.CodeMirror').height($(window).height()-200);
  $('#tutorial-panel').css("min-height", function(){
    return $(window).height()-200;
  });

  //============ render tutorial content from md

  Messenger.options = {
   parentLocations: ['nav'],
  	   // extraClasses: 'messenger-fixed messenger-on-top messenger-on-right',
       theme: 'flat',
     }

     socket.onopen = function(){
      Messenger().post({
        message: 'Server connected.',
        hideAfter : 3,
        type: 'success',
      });
    };

    socket.onerror = function() {
      Messenger().post({
        message: 'Error connecting to server.',
        hideAfter : 3,
        type: 'error',
      });
    };

      // Message received on the socket
      socket.onmessage = function (event) {
          //var line = JSON.parse(event.data);
          // $('#console-output').append(event.data)
          terminal.outputLine(event.data);
        };

      // Send messages.
      var postCode = function () {
        var tutname = $('#tutname').text();
        var content = editor.getValue();
          //socket.send(content);
          //  $.post('/console/sync', 
          //  {
          //     tutname: tutname,
          //     data: content
          // },
          // function(data){
          //     $("console").append(data+"<br>")
          // });

           //============send code via messenger
           Messenger().run({
            action: $.ajax,
            hideAfter : 3,
            successMessage: 'Code sent.',
            errorMessage: 'Error sending code.',
            progressMessage: 'Code on the way...'
          }, {
            /* These options are provided to $.ajax, with success and error wrapped */
            url: '/v1/testbuild/',
            data: {
              tutname:tutname,
              data:content
            },
            type: 'POST',

            error: function(resp){
              if (resp.status === 409)
                return "Error sending code.";
            }
          });

         }

         $('#prev').click(function () {
          getTutorial(prev);
        });
         $('#next').click(function () {
          getTutorial(next);
        });
         $('.nav-item').click(function(){
          getTutorial(this.id.substr(4));
        });

         $('#submitbtn').click(function () {
          postCode();
        });

         $('#clearconsolebtn').click(function () {
         $('#console-output').html("");
  	//$('#console-output').empty();
  });
    //    $('#debugbtn').click(function () {
    // });
});
