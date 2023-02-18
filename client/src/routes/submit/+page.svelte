<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loggedIn, selectedSchools, token } from '../../stores';
	import SearchableTextInput from '../../components/SearchableTextInput.svelte';
	import type { removeDetail, schoolOption, selectDetail } from '../../types';
	import SelectedSchool from '../../components/SelectedSchool.svelte';

	onMount(() => {
		if (!$loggedIn) {
			goto('/login');
		} else {
			if (!$token) {
				goto('/');
			}
		}
	});

	let selectableDataOptions: schoolOption[] = [
		{
			text: "school1",
			attrs: [],
		},
		{
			text: "school2",
			attrs: [],
		},
		{
			text: "school3",
			attrs: [],
		},
	];
	let selected: schoolOption[] = [];
	// Fix this to append to array of selected schools
	// this should trigger pushing onto an array which is passed as props
	// to a component that allows for removing. That means the list will have to be a context
	const handleSelect = (event: { detail: selectDetail; }) => {
		const school = event.detail.option;
		selected = [...selected, school]
		selectableDataOptions = selectableDataOptions.filter( (option) => option.text != event.detail.option.text);
	}

	const handleRemove = (event: {detail: removeDetail; }) => {
		const school = event.detail.option
		selectableDataOptions = [...selectableDataOptions, school];
		selected = selected.filter((selectedSchool) => school.text != selectedSchool.text);
	}
</script>

<div class="flex flex-row">
	<div class="flex flex-col content-between w-6/12">
		{#if selected.length > 0 }			
			{#each selected as school, i }
				<SelectedSchool
					school={school}
					on:remove={handleRemove}
					/>
			{/each}
		{/if}
	</div>
	<div class="flex flex-col items-center w-6/12">
		<h1>Select your schools</h1>
		<SearchableTextInput
		placeholder="Search..."
		data={selectableDataOptions}
		on:select={handleSelect} />
		<div class="mt-40">
			<button>
				Submit Selection
			</button>
		</div>
	</div>
</div>
	