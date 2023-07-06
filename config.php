<?php
namespace QRISApi;

require_once('envparser.php');

use QRISApi\DotEnv;

class DBConnection {
	function __construct() {
		(new DotEnv(__DIR__ . '/.env'))->load();
	}

	function connDB() {
		$host = getenv("DB_HOST");
		$port = getenv("DB_PORT");
		$user = getenv("DB_USER");
		$pass = getenv("DB_PASS");
		$name = getenv("DB_NAME");

		$opt = array(\PDO::ATTR_ERRMODE => \PDO::ERRMODE_EXCEPTION);
		$pdo = new \PDO('mysql:dbname='. $name .';host=' . $host . ';port=' . $port,$user,$pass,$opt);
		return $pdo;
	}
}