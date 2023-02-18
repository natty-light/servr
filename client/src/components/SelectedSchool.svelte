<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { removeDetail, schoolOption, selectDetail } from '../types';
  export let school: schoolOption;
  const dispatch = createEventDispatcher();

  const handleRemoval = (school: schoolOption) => {
    const message: removeDetail = { option: school}
    dispatch("remove", message) 
  }
</script>

<div class="flex flex-row content-between">
  <div>
    {school.text}
  </div>
  <button on:click={(_e) => handleRemoval(school)}>
    x
  </button>
</div> 