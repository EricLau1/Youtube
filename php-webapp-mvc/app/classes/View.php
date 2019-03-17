<?php

namespace app\classes;

class View {

    public function render($view) {
        $view = "../app/views/{$view}.php";
        if(!file_exists($view)) {
            throw new \Exception("Página não existe.");
        }
        return $view;
    }
}