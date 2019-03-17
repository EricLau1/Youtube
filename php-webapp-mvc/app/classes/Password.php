<?php

namespace app\classes;

class Password {

    public static function hash($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    public static function verify($password, $hashedPassword) {
        return password_verify($password, $hashedPassword);
    }
}