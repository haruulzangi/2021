<?php
 
$curl =$_SERVER['HTTP_HOST'].$_SERVER['REQUEST_URI'];

if ($curl != 'screen/admin.php') {
	echo '<!DOCTYPE html>

<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

<title>скрийншот</title>

 
<link rel="stylesheet" href="./css/main.css">
<link href="./css/1.css" rel="stylesheet">

</head>

<body aria-hidden="false"><section id="header" class="f1 fullscreen"><div class="container"><h2><a href="#" target="_blank">админ</a></h2> 
 
 
<p id="message">  зөвхөн дотоод сүлжээнээс хандана уу.  </p>
 
 



</div></section>


<section id="footer"><ul class="copyright"><li>© 2021 Скрийншот</li><li><a href="./feed.php">Feedback</a></li><!--<li><a href="./admin.php">Admin</a></li>--> </ul></section>

</html>';

}
else {
  
	
	
echo	'

<!DOCTYPE html>

<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

<title>скрийншот</title>

 
<link rel="stylesheet" href="./css/main.css">
<link href="./css/1.css" rel="stylesheet">

</head>

<body aria-hidden="false"><section id="header" class="f1 fullscreen"><div class="container"><h2><a href="#" target="_blank">админ</a></h2> 
 
 <p id="message"> HZ{2440a6c77ef002838ca441be4a5ec97d}</p>  
  
 


</div></section>


<section id="footer"><ul class="copyright"><li>© 2021 Скрийншот</li><li><a href="./feed.php">Feedback</a></li><!--<li><a href="./admin.php">Admin</a></li>--> </ul></section>

</html>';
 
 }
?>