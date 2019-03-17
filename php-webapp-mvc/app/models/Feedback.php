<?php

namespace app\models;

class Feedback extends Model {

    protected $table = 'feedbacks';

    public function getFullFeedbacks() {
        $sql = "select u.id as user_id, u.nickname as nickname, u.email as email, a.image as avatar, f.comment as comment
            from {$this->table} as f
            inner join users as u on u.id = f.user
            inner join avatars as a on a.user = u.id order by f.id asc";
        $rs = $this->connection->prepare($sql);
        $rs->execute();
        return $rs->fetchAll();
    }
}