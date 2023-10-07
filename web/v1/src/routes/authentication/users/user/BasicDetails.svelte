<script lang="ts">
    export let user : UserDto
</script>

<div class="flex flex-col gap-y-6 gap-x-4">

    <div class="flex items-center">
        <span class="inline-flex h-[8rem] w-[8rem] flex-shrink-0 items-center justify-center rounded-full bg-gray-500">
          <span class="text-5xl font-medium leading-none text-white">{user.first_name[0]}{user.last_name[0]}</span>
        </span>
        <div class="ml-4">
            <div class="text-3xl font-medium text-gray-900">{user.first_name} {user.last_name}</div>
            <div class="text-lg mt-1 text-gray-500">{user.email}</div>
        </div>
    </div>

    <hr class="border-1 border-gray-200 rounded-full">

    <div class="grid grid-cols-2 gap-y-6 gap-x-4">
        <div class="">
            <label class="text-gray-400">First Name</label>
            <p class="text-2xl">{user.first_name}</p>
        </div>

        <div class="">
            <label class="text-gray-400">Last Name</label>
            <p class="text-2xl">{user.last_name}</p>
        </div>

        <div class="col-span-2">
            <label class="text-gray-400">Email</label>
            <p class="text-2xl">{user.email}</p>
        </div>
    </div>

    <hr class="border-1 border-gray-200 rounded-full">

    <div class="flex flex-col">
        <div>
            <p class="text-xl">Remove User</p>
            <p class="text-sm text-gray-500">This will remove the user from the system</p>
        </div>
        <div class="mt-4 flex">
            <div class="mr-auto py-2 flex">
                <input type="checkbox" class="h-4 w-4 text-red-600 rounded-full border-gray-300 my-auto">
                <label class="ml-2 text-sm text-gray-500 my-auto">I understand the consequences</label>
            </div>
            <button class="h-8 px-4 outline outline-red-600 outline-2 rounded-full text-md font-semibold text-red-600 hover:bg-red-100">Remove</button>
        </div>
    </div>

</div>

