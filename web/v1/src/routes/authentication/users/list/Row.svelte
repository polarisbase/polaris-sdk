<script lang="ts">
    export let index;
    export let user: UserDto
    export let onView: (user: UserDto) => void;
</script>

<tr class="hover:bg-gray-100 hover:cursor-pointer" on:click={() => onView(user)}>
    <td class="whitespace-nowrap py-5 pl-4 pr-3 text-sm sm:pl-4">
        <div class="flex items-center">
            <span class="inline-flex h-10 w-10 items-center justify-center rounded-full bg-gray-500 flex-shrink-0">
                <span class="font-medium leading-none text-white">{user.first_name[0]}{user.last_name[0]}</span>
            </span>
            <div class="ml-4">
                <div class="font-medium text-gray-900">{user.first_name} {user.last_name}</div>
                <div class="mt-1 text-gray-500">{user.email}</div>
            </div>
        </div>
    </td>
<!--    <td class="whitespace-nowrap px-3 py-5 text-sm text-gray-500">-->
<!--        <div class="text-gray-900">Front-end Developer</div>-->
<!--        <div class="mt-1 text-gray-500">Optimization</div>-->
<!--    </td>-->
<!--    <td class="whitespace-nowrap px-3 py-5 text-sm text-gray-500">-->
<!--        <span class="inline-flex items-center rounded-md bg-green-50 px-2 py-1 text-xs font-medium text-green-700 ring-1 ring-inset ring-green-600/20">Active</span>-->
<!--    </td>-->
<!--    <td class="whitespace-nowrap px-3 py-5 text-sm text-gray-500">Member</td>-->
    <td class="relative whitespace-nowrap py-5 pl-3 pr-4 text-right text-sm font-medium sm:pr-4">
<!--        <p on:click={() => onView(user)} class="text-indigo-600 hover:text-indigo-900 hover:cursor-pointer">Edit<span class="sr-only">, {user.first_name} {user.last_name}</span></p>-->
    </td>
</tr>

<!--<tr class="hover:bg-gray-100">-->
<!--    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">{user.first_name}</td>-->
<!--    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">{user.last_name}</td>-->
<!--    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">{user.email}</td>-->
<!--    <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">-->
<!--        <p on:click={() => onView(user)} class="text-indigo-600 hover:text-indigo-900 hover:cursor-pointer">View<span class="sr-only">, {user.first_name} {user.last_name}</span></p>-->
<!--    </td>-->
<!--</tr>-->