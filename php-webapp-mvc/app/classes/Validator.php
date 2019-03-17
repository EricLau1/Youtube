<?php

namespace app\classes;

class Validator {

    public static function validate($params) {
        if(!isArrayEmpty($params)) {
            foreach ($params as $key => $value) {
                $params[$key] = filter_var($value, FILTER_SANITIZE_STRING);
            }
            return $params;
        }
        return null;
    }
}
