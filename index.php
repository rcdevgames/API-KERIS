<?php
namespace QRISApi;

require_once("model.php");

use QRISApi\Model;

$model = new Model();

$res = $model->getUserByIdDevice("3");

var_dump($res);