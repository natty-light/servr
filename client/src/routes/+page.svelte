<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<script lang="ts">
	import { onMount } from 'svelte';
	import { _ping } from './+page';

  onMount( async () => {
    const res = await _ping();
    console.log(res);
  })
</script>
