<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loggedIn, token } from '../../stores';

	onMount(() => {
		if (!$loggedIn) {
			goto('/login');
		} else {
			if (!$token) {
				goto('/');
			}
		}
	});
</script>

<div>
	<h1>Select your schools</h1>
</div>
