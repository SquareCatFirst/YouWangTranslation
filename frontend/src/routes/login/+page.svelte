<script lang="ts">
	import { goto } from '$app/navigation';
	let username = '';
	let password = '';
	let error = '';
	let loading = false;

	function handleLogin() {
		error = '';
		loading = true;

		fetch('/api/v1/auth/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				action: 'POST',
				sets: [],
				orderBy: [],
				page: 0,
				pageSize: 0,
				data: {
					username,
					password
				},
				filter: {},
				authFilter: {}
			})
		})
		.then(response => response.json())
		.then(data => {
			if (data.status === 0) {
				// 登录成功，保存token并跳转
				const userData = data.data[0];
                localStorage.setItem('userId', userData.id);
				//localStorage.setItem('token', userData.token);
				//localStorage.setItem('userId', userData.userId);
				//localStorage.setItem('nickname', userData.nickname);
				//localStorage.setItem('avatarUrl', userData.avatarUrl);
				//localStorage.setItem('role', userData.role);
				//localStorage.setItem('expireAt', userData.expireAt);
				goto('/');
			} else {
				error = data.msg || '登录失败';
			}
		})
		.catch(err => {
			error = '网络错误，请重试';
			console.error(err);
		})
		.finally(() => {
			loading = false;
		});
	}
</script>

<div class="login-container">
	<div class="login-box">
		<h1>登录</h1>
		
		{#if error}
			<div class="error-message">{error}</div>
		{/if}

		<form on:submit|preventDefault={handleLogin}>
			<div class="form-group">
			<label for="username">账号</label>
			<input
				id="username"
				type="text"
				bind:value={username}
				placeholder="请输入用户名/QQ号/邮箱"
					disabled={loading}
				/>
			</div>

			<div class="form-group">
				<label for="password">密码</label>
				<input
					id="password"
					type="password"
					bind:value={password}
					placeholder="请输入密码"
					required
					disabled={loading}
				/>
			</div>

			<button type="submit" disabled={loading}>
				{loading ? '登录中...' : '登录'}
			</button>

			<button type="button" class="qq-login-btn" disabled={loading}>
				QQ登录
			</button>
		</form>

		<p class="register-link">
			暂无账号?QQ登录
		</p>
	</div>
</div>

<style>
	:global(html),
	:global(body) {
		height: 100%;
		margin: 0;
		padding: 0;
	}

	.login-container {
		display: flex;
		justify-content: flex-end;
		align-items: center;
		min-height: 100vh;
		min-width: 800px;
		background: #F6F5F1;
		padding-right: 300px;
	}

	.login-box {
		background: white;
		padding: 40px;
		border-radius: 8px;
		box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
		border: 1px solid #E2E2DD;
		width: 100%;
		max-width: 400px;
	}

	h1 {
		text-align: center;
		color: 2E2E2E;
		margin-bottom: 30px;
		font-size: 28px;
	}

	.error-message {
		background-color: #fee;
		color: #c33;
		padding: 10px 15px;
		border-radius: 4px;
		margin-bottom: 20px;
		border-left: 4px solid #c33;
	}

	.form-group {
		margin-bottom: 20px;
	}

	label {
		display: block;
		margin-bottom: 8px;
		color: #555;
		font-weight: 500;
	}

	input {
		width: 100%;
		padding: 12px;
		border: 1px solid #CFCFC9;
		border-radius: 4px;
		font-size: 16px;
		box-sizing: border-box;
		transition: border-color 0.3s;
	}

	input:focus {
		outline: none;
		border-color: #667eea;
		box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
	}

	input:disabled {
		background-color: #f5f5f5;
		cursor: not-allowed;
	}

	button {
		width: 100%;
		padding: 12px;
		background: #7F8F86;
		color: white;
		border: none;
		border-radius: 4px;
		font-size: 16px;
		font-weight: 600;
		cursor: pointer;
		transition: transform 0.2s, box-shadow 0.2s;
	}

	button:hover:not(:disabled) {
		transform: translateY(-2px);
		box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
	}

	button:disabled {
		opacity: 0.7;
		cursor: not-allowed;
	}

	.qq-login-btn {
		width: 100%;
		padding: 12px;
		background: #E9F0F6;
		color: #3A6EA5;
		border: 1px solid #C7D8E8;
		border-radius: 4px;
		font-size: 16px;
		font-weight: 600;
		cursor: pointer;
		transition: transform 0.2s, box-shadow 0.2s;
		margin-top: 12px;
	}

	.qq-login-btn:hover:not(:disabled) {
		transform: translateY(-2px);
		box-shadow: 0 5px 15px rgba(199, 216, 232, 0.4);
	}

	.qq-login-btn:disabled {
		opacity: 0.7;
		cursor: not-allowed;
	}

	.register-link {
		text-align: center;
		margin-top: 20px;
		color: #666;
	}

	.register-link a {
		color: #667eea;
		text-decoration: none;
		font-weight: 600;
	}

	.register-link a:hover {
		text-decoration: underline;
	}

	@media (max-width: 480px) {
		.login-box {
			padding: 30px 20px;
		}

		h1 {
			font-size: 24px;
		}
	}
</style>
