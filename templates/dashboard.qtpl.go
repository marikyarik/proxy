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
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
    <style>
        html {
            font-size: 14px;
        }
        .btn-circle.btn-xl {
            width: 70px;
            height: 70px;
            padding: 10px 16px;
            border-radius: 35px;
            font-size: 2rem;    
            text-align: center;
        }
        textarea {
            white-space: pre;
            overflow-wrap: normal;
            overflow-x: scroll;
        }
    </style>
</head>
<body>
<div class="container">
   <div class="row mb-3 mt-3 text-center">
      <div class="col-md-12 d-flex justify-content-center">
         <div class="card col-md-4 mb-4 rounded-3 shadow-sm">
            <div class="card-header py-3">
                <h4 class="my-0 fw-normal">Source - Destination Routes</h4>
            </div>
            <div class="card-body">
                <div class="form-group mb-3">
                    <textarea class="header-json form-control" id="routes" rows=5></textarea>
                </div>
                <button class="w-50 btn btn-lg btn-primary" id="save-config">Save</button>
            </div>
         </div>
      </div>
   </div>
   <div class="row mb-3 text-center justify-content-md-center" id="users">
        <div class="col-md-6 d-flex align-items-center justify-content-lg-center" id="add-button">
            <button type="button" class="btn btn-primary btn-circle btn-xl" id="new-btn">+</button>
        </div>
   </div>
</div>
<script>
    var app = app || {};

    app = {
        userUrl: '/proxy-ui/user',
        configUrl: '/proxy-ui/config',
        users: [],
        routes: {},
        timer: 0,
        init: function(data) {
            app.users = data.users;
            app.routes = data.routes;
            app.run();
        },
        run: function() {
            app.render();

            document.getElementById('save-config').onclick = function () {
                if (document.getElementById('routes').value.length > 0 && app.isJson(document.getElementById('routes').value)) {
                    app.routes = JSON.parse(document.getElementById('routes').value);
                } else {
                    app.routes = {};
                }
                const json = JSON.stringify(app.routes)
                if (!app.isJson(json)) {
                    alert('Not valid JSON');
                } else {
                    app.saveConfig(json);
                }
            };

            document.getElementById('new-btn').onclick = function () {
                app.renderNewUser();
            };
        },
        render: function () {
            var html = '';
            for (const key in app.users) {
                app.insertUser(app.users[key], document.getElementById('add-button'));
            }

            document.getElementById('routes').value = Object.keys(app.routes).length !== 0 ? JSON.stringify(app.routes, undefined, 4) : '';
        },
        renderNewUser: function() {
            document.getElementById('add-button').insertAdjacentHTML('beforebegin', app.renderUser());

            [...document.getElementsByClassName('add-btn')].map(item => {
                item.onclick = async function () {
                    var json = this.closest('div').querySelector('textarea').value;
                    if (!app.isValidJson(json)) {
                        alert('Not valid JSON');
                    } else {
                        const user = await app.addUser(json);
                        if (user){
                            app.insertUser(user, document.getElementById('add-button'));
                            this.closest('.user').remove();
                        }
                    }
                }
            });
        },
        insertUser: function(user, div) {
            div.insertAdjacentHTML('beforebegin', app.renderUser(user));

            [...document.getElementsByClassName('delete-btn')].map(item => {
                item.onclick = async function () {
                    var id = this.closest('.card').getAttribute('data-id');
                    const code = await app.deleteUser(id);
                    if (code === 200) {
                        this.closest('.user').remove();
                    }
                }
            });

            [...document.getElementsByClassName('activate-btn')].map(item => {
                item.onclick = function () {
                    var id = this.closest('.card').getAttribute('data-id');
                    app.activateUser(id);
                    location.reload();
                }
            });

            app.prettify();
        },
        renderUser: function(user) {
            const idAttr = user !== undefined ? 'data-id="' + user.id +'"' : '';
            const dataEdit = user !== undefined ? 'true' : 'false';
            const activeText = user !== undefined ? (user.active ? 'Deactivate' : 'Activate') : 'Add';
            const className = user !== undefined && user.active ? 'text-white bg-primary' : '';
            const name = user !== undefined ? user['headers']['X-User-Username'] : "New User";
            const deleteBtn = user !== undefined ? `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger delete-btn">
                                        <button type="button" class="btn-close btn-close-white" aria-label="Delete"></button>
                                    </span>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` : '';
            const json = user !== undefined ? JSON.stringify(user['headers'], undefined, 4) : '';
            const actionClass = user !== undefined ? 'activate-btn' : 'add-btn';

            return `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`<div class="col-md-6 user">
                            <div class="card mb-4 rounded-3 border-primary" `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + idAttr + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`">
                                <div class="card-header py-3 border-primary position-relative `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + className + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`">
                                    <h4 class="my-0 fw-normal">`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + name + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</h4>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + deleteBtn + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`
                                </div>
                                <div class="card-body">
                                    <div class="form-group mb-3">
                                        <textarea class="header-json form-control" data-edit="`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + dataEdit + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`" rows=8>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + json + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</textarea>
                                    </div>
                                    <button class="w-50 btn btn-lg btn-primary `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + actionClass + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`">`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(` + activeText + `)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`</button>
                                </div>
                            </div>
                        </div>`)
//line templates/dashboard.qtpl:1
	qw422016.N().S("`")
//line templates/dashboard.qtpl:1
	qw422016.N().S(`;
        },
        prettify: function() {
            [...document.getElementsByClassName('header-json')].map(item => {
                item.onkeyup = app.delay(async function() {
                    if (this.value.length > 0 && app.isJson(this.value)){
                        this.value = JSON.stringify(JSON.parse(this.value), undefined, 4);

                        if (this.getAttribute('data-edit')) {
                            var id = this.closest('.card').getAttribute('data-id');
                            if (!app.isValidJson(this.value)) {
                                alert('Not valid JSON');
                            } else {
                                const user = await app.editUser(id, this.value);
                                if (user){
                                    app.insertUser(user, this.closest('.user'));
                                    this.closest('.user').remove();
                                }
                            }
                        }
                    }
                }, 1500);               
            });
        },
        delay: function(callback, ms) {
            return function() {
                var context = this, args = arguments;
                clearTimeout(app.timer);
                app.timer = setTimeout(function () {
                    callback.apply(context, args);
                }, ms || 0);
            };
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
                const h = JSON.parse(str);
                if (h['X-User-Username'] == undefined ) {
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
                alert(error.message);
            }
        },
        addUser: async function (json) {
            try {
                const response = await app.request("POST", app.userUrl, json);
                const user = await response.json();
                return user;
            } catch (error) {
                alert(error.message);
            }
            return false;
        },
        editUser: async function (id, json) {
            try {
                const response = await app.request("POST", app.userUrl + "/" + id, json);
                const user = await response.json();
                return user;
            } catch (error) {
                alert(error.message);
            }
            return  false;
        },
        deleteUser: async function (id) {
            try {
                const response = await app.request("DELETE", app.userUrl + "/" + id);
                return response.status;
            } catch (error) {
                alert(error.message);
            }
            return false
        },
        activateUser: async function (id) {
            try {
                await app.request("PATCH", app.userUrl + "/" + id);
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

            return response;
        }
    }
    app.init(`)
//line templates/dashboard.qtpl:273
	qw422016.N().Z(data)
//line templates/dashboard.qtpl:273
	qw422016.N().S(`);
</script>
</body>
</html>
`)
//line templates/dashboard.qtpl:277
}

//line templates/dashboard.qtpl:277
func WriteDashboard(qq422016 qtio422016.Writer, data []byte) {
//line templates/dashboard.qtpl:277
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/dashboard.qtpl:277
	StreamDashboard(qw422016, data)
//line templates/dashboard.qtpl:277
	qt422016.ReleaseWriter(qw422016)
//line templates/dashboard.qtpl:277
}

//line templates/dashboard.qtpl:277
func Dashboard(data []byte) string {
//line templates/dashboard.qtpl:277
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/dashboard.qtpl:277
	WriteDashboard(qb422016, data)
//line templates/dashboard.qtpl:277
	qs422016 := string(qb422016.B)
//line templates/dashboard.qtpl:277
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/dashboard.qtpl:277
	return qs422016
//line templates/dashboard.qtpl:277
}
