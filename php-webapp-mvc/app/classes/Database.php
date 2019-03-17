<?php

namespace app\classes;

use PDO;

class Database {

    public static function connect() {
        $config = require('../config.php');
        $host = $config["db"]["host"];
        $dbname = $config["db"]["name"];
        $db = new PDO("mysql:host={$host};dbname={$dbname};charset=utf8", $config["db"]["user"], $config["db"]["pass"]);
        $db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        return $db;
    }

}