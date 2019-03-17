<?php

require '../bootstrap.php';

use app\classes\Uri;
use app\classes\Router;

$uri = Uri::load();
$routes = require("../app/routes.php");

try {

    $view = new app\classes\View;

    require(Router::load($uri, $routes));

} catch (Exception $e) {
    var_dump($e->getMessage());
}