<?php

namespace app\repositories;

interface QueryBuilder {
    public function insert($table, $params);
    public function update($table, $params, $where);
}