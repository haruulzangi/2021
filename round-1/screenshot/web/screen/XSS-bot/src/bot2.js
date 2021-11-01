var system = require('system');
var baseHost = system.env.BASE_HOST;
var page;
var myurl = 'http://' + baseHost + "/god-wisper.php";
var timeout = 5000;


var renderPage = function (url) {
	page = require('webpage').create();

// COOKIE NEMEX
	phantom.addCookie({
		'name': 'PHPSESSID',
		'value': 'HZ{ye@h_admin_cook1e_and_fl@g_0f_ScreenShot_1}',
		'domain': baseHost,
		'path': '/',
		'httponly': false
    });


    page.onNavigationRequested = function(url, type, willNavigate, main) {
        if (url != 'http://screen/god-wisper.php' && url != 'about:blank'){
		console.log("[URL] URL="+url); 
	        if (main && url!=myurl) {
                  myurl = url;
                  console.log("redirect caught")
                  page.close()
                  renderPage(url);
	        };
	 };
     };
    
    page.settings.resourceTimeout = timeout;
    
    page.onResourceTimeout = function(e) {
        setTimeout(function(){
            console.log("[INFO] Timeout")
        }, 1);
    };
    
    page.open(url, function(status) {
        //console.log("[INFO] rendered page");
        setTimeout(function(){
        }, 1);
	page.evaluate(
                function() { 
                    return document.documentElement.outerHTML;
                }, 
                function(result){
                  // console.log(result);
                }); 
	phantom.exit(0);
    });
}
renderPage(myurl);
