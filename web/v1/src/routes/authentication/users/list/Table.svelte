<script lang="ts">

    import Row from "./Row.svelte";
    import {onMount} from "svelte";
    import SlideOver from "../user/SlideOver.svelte";

    let users = [] as Array<UserDto>;

    let slideOverData = {
        open: false,
        user: {} as UserDto
    }

    $: users = []

    function listUsers() {
        fetch('http://localhost:3000/api/auth/v1/users')
            .then(response => {
                if (!response.ok) {
                    throw new Error(response.statusText)
                }
                return response.json()
            })
            .then(data => {
                users = data.data.users;
            })
    }

    onMount(() => {
        listUsers();
    });

</script>

<SlideOver
    user={slideOverData.user}
    open={slideOverData.open}
    onCloseClick={()=>{slideOverData.open = false}}
    captive={true}
/>

<div class="flex flex-col gap-6">
    <div class="grid grid-cols-2">
        <div class="sm:flex-auto">
            <h1 class="text-3xl font-semibold leading-6 text-gray-900">Users</h1>
            <p class="mt-4 text-lg text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
        </div>
        <div class="flex">
            <button type="button" class="block mt-auto ml-auto rounded-md bg-indigo-600 px-3 py-2 text-center text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Add user</button>
        </div>
    </div>

    <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
            <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                <table class="min-w-full ">
                    <thead class="bg-gray-50">
<!--                    <tr>-->
<!--                        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">First Name</th>-->
<!--                        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Last Name</th>-->
<!--                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Email</th>-->
<!--                        <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">-->
<!--                            <span class="sr-only">Edit</span>-->
<!--                        </th>-->
<!--                    </tr>-->
                    </thead>
                    <tbody class="divide-y divide-gray-200 bg-white">

                    {#each users as user, i}
                        <Row user={user} index={i} onView={clickedUser => {
                            slideOverData.user = clickedUser;
                            slideOverData.open = true;
                        }}/>
                    {/each}

                    <!-- More people... -->
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

