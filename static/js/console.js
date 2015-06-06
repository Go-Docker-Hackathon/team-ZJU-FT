  var socket;
  var url_websocket = "localhost:8081"
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
  var getTutorial = function(tutname){
    $.ajax({
      type: 'GET',
      url: "/v1/tutorial" ,
      data: {
        tutname:tutname
      } ,
      success: function(data){
        //highlight nav bar item
        $('.nav-item').removeClass("active-li");
        $(document.getElementById('nav-'+data.title.toLowerCase())).addClass("active-li");

        document.getElementById('tutorial-title').innerHTML = data.title;
        document.getElementById('page-name').innerHTML = data.title;
        document.getElementById('tutorial-content').innerHTML = marked(data.content);

        editor.setValue(data.code);
        $("#next").html(data.next+" »");
        $("#prev").html("« " +data.prev);

        prev = data.prev;
        next = data.next;
        tutname = data.title; 

      } , 
      dataType: 'json'
    });
  }

  getTutorial(tutname);

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
          $('#console-output').append(event.data)
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
