<?php
namespace QRISApi;

require_once("config.php");

use QRISApi\DBConnection;


class Model {
	public $db;

	function __construct() {
		$this->db = (new DBConnection())->connDB();
	}

	function getUserByIdDevice($id_device) {
		$stm = $this->db->prepare("SELECT * FROM users WHERE device_id = ?");
		$stm->execute([$id_device]);
		return $stm->fetch();
	}

	function insertNewUser($id_device) {
		$stm = "INSERT INTO users (device_id,is_active,last_login,created_date,created_by)"
		$this->db->prepare($stm)->execute([$id_device,1,date('Y-m-d H:i:s'),date('Y-m-d H:i:s'),$device_id]);
		return true;
	}

	function updateUserBy($data) {
		$stm = "UPDATE users SET ";
		foreach ($data as $key => $value) {
			$stm .= $key.'=?,';
		}
		$stm .= " WHERE device_id = ?;";
		$stm = str_replace(", WHERE"," WHERE",$stm);
		$q = $this->db->prepare($stm);
		$params = array_merge($data,[$data["id_device"]]);
		$q->execute($params);
		return true;
	}

	
}
