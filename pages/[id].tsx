import { useEffect, createRef } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import hljs from 'highlight.js'

import styles from '../styles/Viewer.module.css'
import header_styles from '../styles/Header.module.css'
import 'highlight.js/styles/atom-one-dark.css';

import NoteAdd from '@material-ui/icons/NoteAdd'
import { GetServerSideProps } from 'next'


const Viewer = ({ code }: {code: string}) => {

    const codeRef = createRef<HTMLTextAreaElement>();

    const router = useRouter()

    const lines = code.split('\n');

    const html = hljs.highlightAuto(code);

    useEffect(() => {
        if (codeRef.current)
            codeRef.current.innerHTML = html.value;
    }, [html])

    return (
        <div className={styles.container}>
            <Head>
                <title>fastbin</title>
                <meta name="description" content="fastbin: sharing code made faster" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <div className={header_styles["header"]}>
                <div className={header_styles["logo"]}> fastbin </div>
                <div className={header_styles["buttons-container"]}>
                    <div
                        className={header_styles["buttons"]}
                        onClick={() => router.push('/')}
                    ><NoteAdd /></div>
                </div>
            </div>

            <div className={styles['viewer']}>
                <code className={styles["line-numbers"]}>
                    {
                        lines.map((_line, index) => <pre key={index}> {index + 1} </pre>)
                    }
                </code>
                <pre className={styles["code"]}>
                    <code ref={codeRef}>
                    </code>
                </pre>
            </div>
        </div>
    )
}

export const getServerSideProps = (async (context) => {
    const id = context.params?.id
    // const setCode = (data) => {};

    const data = await fetch(`/api/get/${id}`).then(res => res.json())
    if (!data.code) {
        return {
            redirect: {
                destination: '/',
                permanent: false,
            }
        }
    }

    return {
        props: {
            code: data.code,
        }
    }
}) satisfies GetServerSideProps<{ code: string }>

export default Viewer
