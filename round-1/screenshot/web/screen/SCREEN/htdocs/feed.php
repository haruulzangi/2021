<?php
  function escape($str){
      //  $str = str_replace("document", "", $str);
		//$str = str_replace("cookie", "", $str);
		//$str = str_replace("location", "", $str);
	//	$str = str_replace("window", "", $str);
        //$str = str_replace(".", "", $str);
        $str = str_replace("script", "", $str);
	//$str = str_replace("http", "", $str);

       
        return $str;
    }

if($_SERVER['REQUEST_METHOD']=='POST' && isset($_POST['feed']) && !empty($_POST['feed']) ){


    $file_name = 'feed_'.substr(md5(uniqid(rand(), true)),10);
    $fh = fopen("./wishes/$file_name", 'w+');
    if($fh == False){
        die('File write permission is required');
    }
    fwrite($fh, 'feed '."\n".escape($_POST['feed']));
    fclose($fh);

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

<body aria-hidden="false"><section id="header" class="f1 fullscreen"><div class="container"><h2><a href="#" target="_blank">Feedback</a></h2><p id="message">Таны санаа бидэнд үргэлж хэрэгтэй.</p>


<form id="form" class="actions" action="feed.php" method="POST">
 <textarea  name="feed" style="width: 100%; height: 200px"/></textarea>
 

<ul class="actions"><li><button class="button" type="submit">санал явуулах</button></li> </ul>
<p id="message"> <?php if (isset($_POST['feed'])) {
                echo escape($_POST['feed']);
            }
                        ?></p>
</form>


</div></section>

 
<section id="footer"><ul class="copyright"><li>© 2021 Скрийншот</li><li><a href="./feed.php">Feedback</a></li><!--<li>Admin</li>--> </ul></section>


</html>
