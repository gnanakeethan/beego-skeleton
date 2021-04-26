<div class="flex flex-col items-center justify-around content-center"
     style="height:100vh;">
    <div class="flex flex-col items-center"
         style="min-width:300px;max-width:425px;width:80%;min-height: 100px;">
        <div class="font-black text-2xl my-8">{{.AppName}} | {{.Title}}</div>
        <form action="/auth/login?redirect_back={{.RedirectBack}}" method="post"
              class="flex flex-col w-full">
            {{ .xsrfdata }}
            {{if .Flash.errored}}
            <div class="bg-red-200 text-black p-4">
                {{.Flash.error_type}} : {{.Flash.fields.username}}
            </div>
            {{end}}
            <label for="username" class="px-2 pt-4">Username</label>
            <input class="appearance-none block w-full bg-white text-gray-700 border border-gray-200 rounded p-3 leading-tight focus:outline-none focus:bg-gray-200"
                   id="username"
                   name="username"
                   type="text"
                   placeholder="Username">
            <label for="password" class="px-2 pt-4">Password</label>
            <input class="appearance-none block w-full bg-white text-gray-700 border border-gray-200 rounded p-3 leading-tight focus:outline-none focus:bg-gray-200"
                   id="password"
                   name="password"
                   type="password"
                   placeholder="Password">
            <div class="mt-16 m-4 flex flex-row items-center content-center justify-between">
                <a href="/auth/reset" class="">
                    <div class=" bg-yellow-400 px-4 py-2">Reset Password</div>
                </a>
                <button class="bg-green-400 px-4 py-2">Login</button>
            </div>
        </form>
    </div>
</div>