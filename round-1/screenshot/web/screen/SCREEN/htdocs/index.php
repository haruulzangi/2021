<?php
require '/vendor/autoload.php';
use JonnyW\PhantomJs\Client;

if(!isset($_COOKIE['PHPSESSID'])) {
    setcookie('PHPSESSID', md5(rand(1,10000).time()), time() + (86400 * 30), "/"); // 86400 = 1 day
}
if (isset($_COOKIE['PHPSESSID'])){
	$cook= $_COOKIE['PHPSESSID'];
	if (!file_exists("screenshots/".$cook."/")){
	      mkdir("screenshots/".$cook); 
	}

}
$msg='';
if($_SERVER['REQUEST_METHOD']=='GET' && isset($_GET['url']) && !empty($_GET['url']) ){

	

	$client = Client::getInstance();
    $client->getEngine()->addOption('--load-images=true');
    $client->getEngine()->addOption('--ignore-ssl-errors=true');

	$width  = 800;
	$height = 600;
	$top    = 0;
	$left   = 0;

	/** 
	 * @see JonnyW\PhantomJs\Http\CaptureRequest
	 **/
	 $name = substr(md5(uniqid(rand(), true)),10);
   	$request = $client->getMessageFactory()->createCaptureRequest($_GET['url'], 'GET');
	$request->setOutputFile('/var/www/html/screenshots/'.$cook.'/'.$name.'.jpg');
	$request->addSetting('userAgent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063');
	$request->setViewportSize($width, $height);
	$request->setCaptureDimensions($width, $height, $top, $left);

	/** 
	 * @see JonnyW\PhantomJs\Http\Response 
	 **/
	$response = $client->getMessageFactory()->createResponse();
	$client->getEngine()->setPath('/bin/phantomjs');
	// Send the request
	$client->send($request, $response);
	$msg='<i>Таны '.$_GET['url'].' скрийншот </i><b><a href="screenshots/'.$cook.'/'.$name.'.jpg'.'"  target="_blank">энд</a> </b> <i>хууладлаа</i>';


}





?>

<!DOCTYPE html>

<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

<title>скрийншот</title>

 
<link rel="stylesheet" href="./css/main.css">
<link href="./css/1.css" rel="stylesheet">

</head>

<body aria-hidden="false"><section id="header" class="f1 fullscreen"><div class="container"><h2><a href="#" target="_blank">Скрийншот</a></h2><p id="message">Вэб скрийншот авах онлайн үйлчилгээ</p>


<form id="form" class="actions" action="index.php" method="GET"><input type="text" name="url" id="url" placeholder="URL ээ оруулна уу"><br> 

<ul class="actions"><li><button class="button" type="submit">скрийншот авах</button></li> </ul>
<p id="message"><?php echo $msg; ?></p>
</form>


</div></section>

 
<section id="footer"><ul class="copyright"><li>© 2021 Скрийншот</li><li><a href="./feed.php">Feedback</a></li><!--<li><a href="./admin.php">Admin</a></li>--> </ul></section>


</html>