<?php

namespace app\models;
use app\classes\Database;
use app\models\Persist;

abstract class Model {

    public $connection;

    // HAHA
    public function __construct() {
        $this->connection = Database::connect();
    }

    public function getAll() {
        $sql = "select * from {$this->table}";
        $rs = $this->connection->prepare($sql);
        $rs->execute();
        return $rs->fetchAll();
    }

    public function findOne(array $params) {
        $sql = "select * from {$this->table} where {$params[0]} = ?";
        $rs = $this->connection->prepare($sql);
        $rs->bindValue(1, $params[1]);
        $rs->execute();
        return $rs->fetch();
    }

    public function create($entity) {
        $sql = Persist::insert($this->table, $entity);
        $rs = $this->connection->prepare($sql);
        return $rs->execute($entity);
    }

    public function update($entity, $where) {
        $sql = Persist::update($this->table, $entity, $where);
        $rs = $this->connection->prepare($sql);
        return $rs->execute($entity); 
    }
}