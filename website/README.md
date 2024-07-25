This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Requirements

`nix` installed on your system, or `bun` installed globally if you prefer to not use `nix`

## Getting Started

Enter Nix shell environment

```
$ nix-shell
```

Install dependencies

```
[nix-shell]$ bun i
```

Run the development server:

```bash
[nix-shell]$ bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying any file. The page auto-updates as you edit the file.

## Deployment

Pushes to the `main` branch are automatically deployed to Cloudflare Pages. The
deployment dashboard can be found here:
https://dash.cloudflare.com/8a3fd12f290cd3ecfdacccbbebf096c6/pages/view/kardinal-landing-page

To ensure your code will build when pushed, build the entire app locally to ensure
there are no typescript errors (currently the dev server only lazy-compiles
pages that you load). This experience can be improved in the future.

```
[nix-shell]$ bun run build
```

If there are any errors, fix them before pushing. You can also auto-fix some
linter errors:

```
[nix-shell]$ bun run lint:fix
```

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy
