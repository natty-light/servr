<script lang="ts">
	import { goto } from '$app/navigation';
	import { _login } from './+page';
	import { loggedIn, token } from '../../stores';
	let username: string, password: string;

	const handleLogin = async () => {
		const destructuredToken = await _login(username, password);

		if (destructuredToken) {
			token.set(destructuredToken);
			loggedIn.set(true);
			goto('/submit');
		} else {
			window.alert('Unable to log in');
		}
	};
</script>

<div class="flex flex-col items-center">
	<h1 class="text-2xl">
    Login Form
  </h1>
	<input 
    class="border-2 m-2"
    type="text"
    bind:value={username}
    required={true} />
	<input
    class="border-2 m-2"
    type="password"
    bind:value={password}
    required={true} />
	<button
    class="border-2 rounded-sm" 
    on:click={handleLogin}
    >
    submit 
  </button>
</div>
