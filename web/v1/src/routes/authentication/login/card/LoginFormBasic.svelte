<script lang="ts">
    let emailAddress:string = '';
    let password:string = '';
    function login(){
        console.log(emailAddress);
        console.log(password);
        const res = fetch('http://localhost:3000/api/auth/v1/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: emailAddress,
                password: password
            }),
        });
        res.then((res)=>{
            console.log(res);
            if(res.status === 200){
                console.log('success');
            }else if(res.status === 401){
                console.log('unauthorized');
            }else{
                console.log('error');
            }
        });
    }
</script>

<div class="
xl:w-1/3
lg:w-2/5
md:w-1/2
sm:w-11/12 sm:h-auto sm:my-auto sm:shadow-lg sm:border-none
w-full h-full mx-auto overflow-hidden rounded-lg bg-white

">
    <div class="px-4 py-5 sm:p-6">
        <div class="flex flex-col gap-2">
            <div class="pb-2">
                <p>Login with your email address and password</p>
            </div>
            <div>
                <label for="email-address" class="block text-sm font-medium text-gray-700">Email address</label>
                <div class="mt-1">
                    <input type="email" id="email-address" name="email-address" autocomplete="email" bind:value={emailAddress}
                           class="block w-full h-12 px-2 rounded-md border-2 border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 text-lg">
                </div>
            </div>
            <div>
                <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                <div class="mt-1">
                    <input type="password" id="password" name="password" bind:value={password}
                           class="block w-full h-12 px-2 rounded-md border-2 border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 text-lg">
                </div>
            </div>
            <div class="pt-2">
                <button class="block w-full h-12 px-2 rounded-md border-2 border-gray-300 hover:bg-gray-600 hover:text-white" on:click={()=>{
                    login();
                }}>Login</button>
            </div>
        </div>
    </div>
</div>
