/*
Navicat MySQL Data Transfer

Source Server         : aaa
Source Server Version : 50547
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50547
File Encoding         : 65001

Date: 2024-03-27 01:20:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for web_components
-- ----------------------------
DROP TABLE IF EXISTS `web_components`;
CREATE TABLE `web_components` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(65) DEFAULT '' COMMENT '名称',
  `code` text COMMENT '代码',
  `state` tinyint(4) DEFAULT '0' COMMENT '状态',
  `price` bigint(20) unsigned DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL COMMENT '代码',
  PRIMARY KEY (`id`),
  KEY `idx_web_components_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of web_components
-- ----------------------------
INSERT INTO `web_components` VALUES ('1', '2024-03-20 00:41:18', '2024-03-27 01:18:46', null, 'qlrender.vue', '     <template>\n        <template v-for=\"item in dome\">\n              <qlcustom  :content=\"item.type\" :layer=\"item\" :vdata=\"dd\" />\n        </template>\n       {{dd}}\n     </template>\n    <script setup>\n        import {ref,reactive,defineProps,getCurrentInstance,computed} from \'vue\';\n        const {proxy} =getCurrentInstance();\n        const dome = reactive([\n            {\n              id:2,\n              alias:\"fafafa\",\n              type:\"qlinput.vue\",\n              style:{width:\'150px\'},\n              data:\"1111\",\n              child:[{\n                id:3,\n                type:\"qlinput.vue\",\n                data:false,\n              }]\n            },\n           {\n              id:3,\n              alias:\"form\",\n              type:\"qlform.vue\",\n              style:{width:\'150px\'},\n              bind:{size:\'mini\'},\n              data:\"1111\",\n              child:[{\n                id:3,\n                type:\"qlinput.vue\",\n                data:false,\n              }]\n            }\n          ])\n            const dd=ref({})\n            dome.forEach((e)=>{\n              dd[e.alias]=e.data\n            })\n        const aaaa=()=>{\n            proxy.$get(\"aaa/bbb\")\n        }\n    </script>\n    ', '0', '0', '测试渲染');
INSERT INTO `web_components` VALUES ('2', '2024-03-21 00:46:57', '2024-03-24 23:00:30', null, 'base.vue', '<template>\n  <h2>输入代码111111</h2>\n  <el-button>按钮</el-button>\n  <qlcustom content=\"qlinput.vue\" />\n</template>\n <script setup>\n    import {ref,defineProps,getCurrentInstance} from \'vue\';\n    const {proxy} =getCurrentInstance();\n    const a=ref(\"测试22\")\n    const aa=ref(\"测试22\")\n\n    const aaaa=()=>{\n        proxy.$get(\"aaa/bbb\")\n    }\n</script>', '1', '0', null);
INSERT INTO `web_components` VALUES ('3', '0000-00-00 00:00:00', '2024-03-21 23:49:39', '2024-03-21 23:54:44', '我eee2121', '     <template>\n        输入代码\n     </template>\n    <script setup>\n        import {ref,defineProps,getCurrentInstance} from \'vue\';\n        const {proxy} =getCurrentInstance();\n        const a=ref(\"测试22\")\n        const aa=ref(\"测试22\")\n\n        const aaaa=()=>{\n            proxy.$get(\"aaa/bbb\")\n        }\n    <//script>\n    ', '0', '0', null);
INSERT INTO `web_components` VALUES ('4', '0000-00-00 00:00:00', '2024-03-26 23:52:33', null, 'qlinput.vue', '     <template>\n        <el-input v-model.trim=\"data\" :style=\"prop.layer.style\" />\n     </template>\n    <script setup>\n        import {ref,defineProps,defineEmits,getCurrentInstance,computed,defineModel} from \'vue\';\n        const {proxy} =getCurrentInstance();\n        const prop=defineProps({\n          layer:Object,\n          data:{type:String,default:\"\"},\n        })\n        //const data = defineModel(\'data\')\n        const emit=defineEmits([\'update:data\'])\n        const data=computed({\n            get:()=>{\n                return prop.data\n            },\n            set:(val)=>{\n                emit(\'update:data\',val)\n            }\n        })\n        const aaaa=()=>{\n            proxy.$get(\"aaa/bbb\")\n        }\n    </script>\n    ', '0', '0', null);
INSERT INTO `web_components` VALUES ('5', '2024-03-21 23:56:38', '2024-03-21 23:56:49', '2024-03-22 23:15:22', '哈哈哈', '     <template>\n        输入代码\n     </template>\n    <script setup>\n        import {ref,defineProps,getCurrentInstance} from \'vue\';\n        const {proxy} =getCurrentInstance();\n        const a=ref(\"测试22\")\n        const aa=ref(\"测试22\")\n\n        const aaaa=()=>{\n            proxy.$get(\"aaa/bbb\")\n        }\n    <//script>\n    ', '1', '0', null);
INSERT INTO `web_components` VALUES ('6', '2024-03-20 00:41:18', '2024-03-27 01:10:29', null, 'qlform.vue', '     <template>\n        <el-form ref=\"formRef\" >\n           <el-form-item v-for=\"item in layer.child\">\n              <qlcustom  :content=\"item.type\" :layer=\"item\" :vdata=\"dd\" />\n            </el-form-item>\n        </el-form>\n     </template>\n    <script setup>\n      import {ref,defineProps,defineEmits,getCurrentInstance,computed,defineModel} from \'vue\';\n      const {proxy} =getCurrentInstance();\n      const prop=defineProps({\n        layer:Object,\n        data:{type:String,default:\"\"},\n      })\n\n        const aaaa=()=>{\n            proxy.$get(\"aaa/bbb\")\n        }\n    </script>\n    ', '0', '0', '');
