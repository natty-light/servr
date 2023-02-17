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

<div>
	<h1>Login Form</h1>
	<input type="text" bind:value={username} required={true} />
	<input type="password" bind:value={password} required={true} />
	<button on:click={handleLogin}> submit </button>
</div>
