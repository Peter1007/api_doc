$(document).ready(function(){
    api_desc.service_url = '/beta_' + (/http:\/\/[^\.]+\./.exec(api_desc.service_url)[0].slice(7, -1)) + (/\.com\/.+/.exec(api_desc.service_url)[0].substring(4));
    var $light = $('pre').eq(0), $show = $('pre').eq(1);

    $light.after('<textarea id="edit-code" style="display: none; width:95%; margin:13px 10px; padding: 13px; box-sizing: content-box;"></textarea><div style="padding: 15px 0; text-align: right;"><select id="test-platform" style="margin: 0; width: 100px;"><option value="/beta_">Beta</option><option value="/gamma_">Gamma</option></select><button class="btn btn-primary" id="btn-light" style="display: none; margin-left: 15px;" type="button">高亮代码</button><button class="btn btn-success" id="btn-run" style="margin-left: 15px;" type="button">发送请求</button></div>');
    var loading = '<img style="display: block; margin: auto; height: 150px;" src="/img/loading.gif">';
    var $edit = $('#edit-code'), $platform = $('test-platform'), $btn_light = $('#btn-light'), $btn_run = $('#btn-run'), $loading = $('#test-loading');

    $edit.val(JSON.stringify(api_desc.request, null, 4));

    $light.on('dblclick', function(e){
        e.preventDefault();
        $edit.height($light.height()).css('display', 'block');
        $btn_light.show();
        $light.hide();
    });

    $btn_light.on('click', function(e){
        e.preventDefault();
        $light.html($edit.val()).show();
        $btn_light.hide();
        $edit.hide();
        prettyPrint();
    });

    $('#test-platform').on('change', function(){
        api_desc.service_url = api_desc.service_url.replace(/(^\/beta_|^\/gamma_)/, $(this).val())
    });

    $btn_run.on('click', function(){
        $show.html(loading);
        $.ajax({
            url: api_desc.service_url,
            type: 'post',
            data: $edit.val().replace(/\r/g, ''),
            success: function(data){
                if(!data){
                    $show.html('无返回值');
                    return;
                }
                if(typeof(data) == 'string'){
                    $show.html(JSON.stringify(JSON.parse(data), null, 4));
                    prettyPrint();
                }else{
                    $show.html(JSON.stringify(data, null, 4));
                    prettyPrint();
                }
            }
        })
    });
});
    