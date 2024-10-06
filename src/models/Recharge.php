<?php

class Recharge {
    private static $conn;

    public static function init($conn) {
        self::$conn = $conn;
    }

    public static function add_recharge($volume, $re_date, $exp_date) {

        try {
            die(self::$conn);
            $stmt = self::$conn->prepare("INSERT INTO recharge (
                $volume,
                $re_date,
                $exp_date
             ) VALUES (?, ?, ?)");

            $stmt->bindParam(1, $volume);
            $stmt->bindParam(2, $re_date);
            $stmt->bindParam(3, $exp_date);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function update_recharge($id, $name_updated) {
        try {
            $stmt = self::$conn->prepare("UPDATE recharge set name = $name_updated WHERE customer.id = $id");
            $stmt->bindParam(1, $name_updated);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function delete_recharge($id) {
        try {
            $stmt = self::$conn->prepare("DELETE recharge WHERE customer.id = $id");
            $stmt->bindParam(1, $id);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }
}