  var name="redis"
  var editor

  function getUrlParameter(sParam)
  {
      var sPageURL = window.location.search.substring(1);
      var sURLVariables = sPageURL.split('&');
      for (var i = 0; i < sURLVariables.length; i++) 
      {
          var sParameterName = sURLVariables[i].split('=');
          if (sParameterName[0] == sParam) 
          {
              return sParameterName[1];
          }
      }
  }

  $(document).ready(function () {
      $("#nav-templates").click();
        editor = CodeMirror.fromTextArea(document.getElementById("code"), {
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

    dockerfile = getUrlParameter('tutname');

          var getTemplate = function(name){
    $.ajax({
      type: 'GET',
      url: "/v1/name" ,
      data: {
        name :name
      } ,
      success: function(data){
        $('.nav-item-temp').removeClass("active-li");
        // $(document.getElementById('nav-'+data.title.toLowerCase())).addClass("active-li");
        $(document.getElementById('nav-'+name)).addClass("active-li");
        document.getElementById('template-title').innerHTML = name;
        console.log(data.dockerfile);
        editor.setValue(data.dockerfile);
      } , 
      dataType: 'json'
    });
  }

  getTemplate(name);

         $('.nav-item').click(function(){
          window.location.href='./page?tutname='+this.id.substr(4);
         });

         $('.nav-item-temp').click(function(){
          getTemplate(this.id.substr(4));
        });



}); 