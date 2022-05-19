// Code generated by qtc from "dashboard.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/dashboard.qtpl:1
package templates

//line templates/dashboard.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/dashboard.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/dashboard.qtpl:1
func StreamDashboard(qw422016 *qt422016.Writer, data []byte) {
//line templates/dashboard.qtpl:1
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <style>
        body {
            background: #B24592;
            background: linear-gradient(to left, #B24592, #F15F79);
            font-family: arial;
            font-size: 14px;
        }

        html {
            box-sizing: border-box;
        }

        .row {
            width: 100%;
            display: inline-block;
            flex-wrap: wrap;
            box-sizing: border-box;
        }
        .box {
            width: 20%;
            float: left;
            position: relative;
            border-radius: 3px;
            background: rgba(0,0,0,0.2);
            border: 5px solid #c90076;
            margin: 20px;
            box-shadow: 0 1px 1px rgba(0,0,0,0.1);
        }
        .box-header {
            color: #444;
            display: block;
            padding: 10px;
            position: relative;
            border-bottom: 1px solid #c90076;
        }
        .box-title {
            display: inline-block;
            font-size: 18px;
            margin: 0;
            line-height: 1;
        }
        .box-content {
            border-top-left-radius: 0;
            border-top-right-radius: 0;
            border-bottom-right-radius: 3px;
            border-bottom-left-radius: 3px;
            padding: 10px;
        }
        .box-content textarea, .box-content input {
            width: 100%;
            font-size: 14px;
        }
        .box-content textarea {
            height: 20vh;
        }
        .btn {
            border-radius: 3px;
            -webkit-box-shadow: none;
            box-shadow: none;
            border: 1px solid #c90076;
            background-color: #c90076;
            cursor: pointer;
            float:right;
            font-size: 14px;
        }
        .btn.disabled {
            background-color: #c0c0c0;
        }
    </style>
</head>
<body>
<div class="row">
    <div class="box">
        <div class="box-header">
            <h3 class="box-title">Edit Proxy Url</h3>
        </div>
        <div class="box-content">
            <input id="proxy" value=""/>
            <button class="btn" id="save-config">Save</button>
        </div>
    </div>
    <div class="box">
        <div class="box-header">
            <h3 class="box-title">Add User</h3>
        </div>
        <div class="box-content">
            <textarea></textarea>
            <button class="btn" id="add">Add</button>
        </div>
    </div>
</div>
<div class="row" id="users">
</div>
<script>
    var app = app || {};

    app = {
        userUrl: '/proxy-ui/user',
        configUrl: '/proxy-ui/config',
        users: [],
        config: {},
        init: function(data) {
            app.users = data.users;
            app.config = data;
            delete app.config.users;
            app.run();
        },
        run: function() {
            app.render();

            document.getElementById('save-config').onclick = function () {
                app.config.proxy_url = document.getElementById('proxy').value;
                const json = JSON.stringify(app.config)
                if (!app.isJson(json)) {
                    alert('Not valid JSON');
                } else {
                    app.saveConfig(json);
                }
            };

            document.getElementById('add').onclick = function () {
                var json = this.closest('div').querySelector('textarea').value;
                if (!app.isValidJson(json)) {
                    alert('Not valid JSON');
                } else {
                    app.addUser(json);
                }
            };

            [...document.getElementsByClassName('edit')].map(item => {
                item.onclick = function () {
                    var json = this.closest('div').querySelector('textarea').value;
                    var id = this.closest('.collapse-parent').getAttribute('data-id');
                    if (!app.isValidJson(json)) {
                        alert('Not valid JSON');
                    } else {
                        app.editUser(id, json);
                    }
                }
            });

            [...document.getElementsByClassName('delete')].map(item => {
                item.onclick = function () {
                    var id = this.closest('.collapse-parent').getAttribute('data-id');
                    app.deleteUser(id);
                }
            });

            [...document.getElementsByClassName('activate')].map(item => {
                item.onclick = function () {
                    var id = this.closest('.collapse-parent').getAttribute('data-id');
                    app.users[id].active = !app.users[id].active;
                    const json = JSON.stringify(app.users[id])
                    if (!app.isValidJson(json)) {
                        alert('Not valid JSON');
                    } else {
                        app.editUser(id, json);
                    }
                }
            });

            [...document.getElementsByClassName('collapse')].map(item => {
                item.onclick = function () {
                    var x = this.closest('.collapse-parent').querySelector('.collapse-body');
                    if (x.style.display === "none") {
                        x.style.display = "block";
                    } else {
                        x.style.display = "none";
                    }
                }
            });
        },
        render: function () {
            var html = '';
            for (const key in app.users) {
                const active = app.users[key].active ? 'Active' : 'Not active';
                const className = app.users[key].active ? '' : 'disabled';
                html += `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`<div class="box collapse-parent" data-id="`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + key + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`">
                            <div class="box-header">
                                <h3 class="box-title">`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + app.users[key].name + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</h3>
                                <button class="btn activate `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + className + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`">`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + active + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</button>
                                <button class="btn collapse">Edit</button>
                                <button class="btn delete">Delete</button>
                            </div>
                            <div class="box-content collapse-body" style="display: none">
                                <textarea>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + JSON.stringify(app.users[key]) + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</textarea>
                                <button class="btn edit">Save</button>
                            </div>
                        </div>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`;
            } 
            document.getElementById('users').insertAdjacentHTML('beforeend', html);

            document.getElementById('proxy').value = app.config.proxy_url;
        },
        isJson: function  (str) {
            try {
                const u = JSON.parse(str);
            } catch (e) {
                return false;
            }
            return true;
        },
        isValidJson: function  (str) {
            try {
                const u = JSON.parse(str);
                if (u.name == undefined || u.active == undefined) {
                    return false;
                }
            } catch (e) {
                return false;
            }
            return true;
        },
        saveConfig: async function (json) {
            try {
                await app.request("POST", app.configUrl, json);
            } catch (error) {
                console.log(error.message);
            }
        },
        addUser: async function (json) {
            try {
                await app.request("POST", app.userUrl, json);
            } catch (error) {
                console.log(error.message);
            }
        },
        editUser: async function (id, json) {
            try {
                await app.request("POST", app.userUrl + "/" + id, json);
            } catch (error) {
                console.log(error.message);
            }
        },
        deleteUser: async function (id) {
            try {
                await app.request("DELETE", app.userUrl + "/" + id);
            } catch (error) {
                console.log(error.message);
            }
        },
        request: async function (method, url, data) {
            const fetchOptions = {
                method: method,
                headers: {
                    "Content-Type": "application/json",
                    Accept: "application/json"
                }
            };

            if ("POST" === method) {
                fetchOptions.body = data
            }

            const response = await fetch(url, fetchOptions);
            if (!response.ok) {
                const jsonErr = await response.json();
                throw new Error(jsonErr.error);
            }

            location.reload();
        }
    }
    app.init(`)
//line templates/dashboard.qtpl:268
	qw422016.N().Z(data)
//line templates/dashboard.qtpl:268
	qw422016.N().S(`);
</script>
</body>
</html>
`)
//line templates/dashboard.qtpl:272
}

//line templates/dashboard.qtpl:272
func WriteDashboard(qq422016 qtio422016.Writer, data []byte) {
//line templates/dashboard.qtpl:272
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/dashboard.qtpl:272
	StreamDashboard(qw422016, data)
//line templates/dashboard.qtpl:272
	qt422016.ReleaseWriter(qw422016)
//line templates/dashboard.qtpl:272
}

//line templates/dashboard.qtpl:272
func Dashboard(data []byte) string {
//line templates/dashboard.qtpl:272
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/dashboard.qtpl:272
	WriteDashboard(qb422016, data)
//line templates/dashboard.qtpl:272
	qs422016 := string(qb422016.B)
//line templates/dashboard.qtpl:272
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/dashboard.qtpl:272
	return qs422016
//line templates/dashboard.qtpl:272
}
