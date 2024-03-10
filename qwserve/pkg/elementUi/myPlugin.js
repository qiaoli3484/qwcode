const myPlugin = {
    install(app, options) {
        // 配置此应用
        app.config.globalProperties.layer_show = function (title, url, w, h) {
            if (title == null || title == '') {
                title = false;
            };
            if (url == null || url == '') {
                url = "404.html";
            };
            if (w == null || w == '') {
                w = window.innerWidth * 0.9;
            };
            if (h == null || h == '') {
                h = ($(window).height() - 50);
            };
            layer.open({
                type: 2,
                area: [w + 'px', h + 'px'],
                fix: false, //不固定
                maxmin: true,
                shade: 0.4,
                resize: true,
                title: title,
                content: url,
                //closeBtn: 0
            });
        }


        app.config.globalProperties.del_sure = function (content) {
            return this.$confirm('删除 ' + content, '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
        }

        app.config.globalProperties.$post = function (url, data) {
            return new Promise((resolve, reject) => {
                axios.post(url, data, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                }).then(res => {
                    if (res.data.code == 200) {
                        resolve(res.data);
                    } else {
                        this.$message({
                            showClose: true,
                            message: res.data.msg,
                            type: 'error'
                        });
                        reject()
                    }
                }).catch(() => {
                    this.$message({
                        showClose: true,
                        message: '服务器消失，请联系管理等待...',
                        type: 'warning'
                    });
                    reject()
                })

            })
        }
        app.config.globalProperties.$get = function (url, data) {
            return new Promise((resolve, reject) => {
                axios.get(url, {
                    params: data
                }).then(res => {
                    if (res.data.code == 200) {
                        resolve(res.data);
                    } else {
                        this.$message({
                            showClose: true,
                            message: res.data.msg,
                            type: 'error'
                        });
                        reject()
                    }
                }).catch(() => {
                    this.$message({
                        showClose: true,
                        message: '服务器消失，请联系管理等待...',
                        type: 'warning'
                    });
                    reject()
                })

            })
        }
        app.config.globalProperties.$del = function (url, data) {
            return new Promise((resolve, reject) => {
                axios.delete(url, {
                    params: data
                }).then(res => {
                    if (res.data.code == 200) {
                        resolve(res.data);
                    } else {
                        this.$message({
                            showClose: true,
                            message: res.data.msg,
                            type: 'error'
                        });
                        reject()
                    }
                }).catch(() => {
                    this.$message({
                        showClose: true,
                        message: '服务器消失，请联系管理等待...',
                        type: 'warning'
                    });
                    reject()
                })

            })
        }
        app.config.globalProperties.$put = function (url, data) {
            return new Promise((resolve, reject) => {
                axios.put(url, data, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                }).then(res => {
                    if (res.data.code == 200) {
                        resolve(res.data);
                    } else {
                        this.$message({
                            showClose: true,
                            message: res.data.msg,
                            type: 'error'
                        });
                        reject()
                    }
                }).catch(() => {
                    this.$message({
                        showClose: true,
                        message: '服务器消失，请联系管理等待...',
                        type: 'warning'
                    });
                    reject()
                })

            })
        }
        /*
        app.config.globalProperties.del_sure = function(str, callback) {
            this.$confirm('永久删除->'+str+'<-, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                callback();
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            });
        }
        */

        app.config.globalProperties.closedown = function () {
            //window.parent.wv.更新表格();
            let index = parent.layer.getFrameIndex(window.name);
            parent.layer.close(index);
        }
    }
}