<?php

namespace app\models;
use app\repositories\QueryBuilder;

class Persist implements QueryBuilder {

    public function insert($table, $params) {
        $sql = "insert into {$table} (";
        $sql .= implode(',', array_keys($params)) . ") values (";
        $sql .= ":". implode(',:', array_keys($params)) . ")";
        return $sql;
    } 

    public function update($table, $params, $where) {
        $sql = "update {$table} set ";
        unset($params["id"]);
        $sql .= implode(",", array_map(function ($param) {
            return "{$param} = :{$param}";
        }, array_keys($params)));
        $sql .= " where {$where[0]} = :{$where[0]}";
        return $sql;
    }
}