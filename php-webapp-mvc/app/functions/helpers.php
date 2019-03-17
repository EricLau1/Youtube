<?php

function toJson($data) {
    header('Content-type: application/json');
    return json_encode($data);
}

function redirect($route = '/') {
    header("location: {$route}");
}

function dd($dump) {
    var_dump($dump);
    die();
}

function isArrayEmpty($attributes) {
    if(!is_array($attributes)) {
        throw new Exception("É necessário passar um Array como parâmentro.");
    }
    foreach($attributes as $attr => $value) {
        if(empty($value)) {
            return true;
        }
    }
    return false; // array não está vazio...
}

function isAuth() {
    if(isset($_SESSION['AUTH'])) {
        $auth = $_SESSION['AUTH'];
        return $auth['is_valid'];
    }
    return false;
}

function flash() {
    $message = '';
    if(isset($_SESSION['MESSAGE'])) {
        $message = $_SESSION['MESSAGE'];
        unset($_SESSION['MESSAGE']);
    }
    return $message;
}