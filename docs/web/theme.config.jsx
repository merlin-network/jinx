import { useConfig } from 'nextra-theme-docs'
import { useRouter } from 'next/router'

export default {
    logo: <span>Jinx</span>,
    logoLink: '/',
    project: {
        link: 'https://github.com/berachain/jinx',
    },
    docsRepositoryBase: "https://github.com/berachain/jinx",
    banner: {
        key: '2.0-release',
        text: <a href="https://medium.com/berachain-foundation/introducing-jinx-vm-2a0b77d777f8" target="_blank">
          🎉 Introducing Jinx Ethereum! 
        </a>,
    },
    useNextSeoProps() {
        const { route } = useRouter()
        if (route !== '/') {
            return {
                titleTemplate: '%s – Jinx Ethereum Docs'
            }
        }
    },
    head: function useHead() {
        const { title } = useConfig()
        const socialCard = '/header.png'
        return (
            <>
                <meta name="msapplication-TileColor" content="#fff" />
                <meta name="theme-color" content="#fff" />
                <meta name="viewport" content="width=device-width, initial-scale=1.0" />
                <meta httpEquiv="Content-Language" content="en" />
                <meta
                    name="description"
                    content="Jinx Ethereum brings EVM to Cosmos in a new way"
                />
                <meta
                    name="og:description"
                    content="Jinx Ethereum brings EVM to Cosmos in a new way"
                />
                <meta name="twitter:card" content="summary_large_image" />
                <meta name="twitter:image" content="/header.png" />
                <meta name="twitter:site:domain" content="https://jinx.berachain.dev/" />
                <meta property="twitter:description" content="Jinx Ethereum brings EVM to Cosmos in a new way"/>
                <meta name="twitter:url" content="https://jinx.berachain.dev/" />
                <meta
                    name="og:title"
                    content={title ? title + ' – Jinx Ethereum' : 'Jinx Ethereum'}
                />
                <meta name="og:image" content={socialCard} />
                <meta name="apple-mobile-web-app-title" content="Jinx Ethereum" />
                <link rel="icon" href="/milky-way.png" type="image/png" />
                <link rel="icon" href="/milky-way.ico"/>
                <link
                    rel="icon"
                    href="/berachain.svg"
                    type="image/svg+xml"
                    media="(prefers-color-scheme: dark)"
                />
                <link
                    rel="icon"
                    href="/berachain.png"
                    type="image/png"
                    media="(prefers-color-scheme: dark)"
                />
            </>
        )
    },
    editLink: false,
    feedback: false,
    sidebar: {
        titleComponent({ title, type }) {
            if (type === 'separator') {
                return <span className="cursor-default">{title}</span>
            }
            return <>{title}</>
        },
        defaultMenuCollapseLevel: 1,
        toggleButton: true,
    },
    footer: {
        text: (
            <div>
                <p>
                    © {new Date().getFullYear()} Blackchain Foundation.
                </p>
            </div>
        )
    },
    gitTimestamp: false,
}