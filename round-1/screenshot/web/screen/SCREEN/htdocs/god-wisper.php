<?php
if($_COOKIE["PHPSESSID"]=='HZ{ye@h_admin_cook1e_and_fl@g_0f_ScreenShot_1}'){

$dh = opendir('./wishes');
?>

<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<body><table>
                <?php
                    while($file = readdir($dh)){
                        if($file != "." && $file != "..") {

                            $content = file("./wishes/$file");
                            echo "<tr><td>$content[0]</td>\n";
                            echo "\n";
                            echo "<td>".$content[1]."</td>";
                            echo "</tr>";
							unlink("./wishes/$file"); // for log only
							//fwrite($fh, "admin". "\n"." шалгасан.");
							//fclose($fh);
							break;
                        }
						}
} else { die();}               ?>
</table>
</body>
</html>
