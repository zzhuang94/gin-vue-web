-- 机器表增加 health 字段：0 故障，1 正常
ALTER TABLE `machine` ADD COLUMN `health` INT NOT NULL DEFAULT 1 COMMENT '健康状态 0故障 1正常' AFTER `name`;
