import { useEffect, createRef } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import hljs from 'highlight.js'

import styles from '../styles/Viewer.module.css'
import header_styles from '../styles/Header.module.css'
import 'highlight.js/styles/atom-one-dark.css';

import NoteAdd from '@material-ui/icons/NoteAdd'
import { GetServerSidePropsContext } from 'next'
import { getData } from './api/get/[id]'


const Viewer = ({ code }: { code: string }) => {

    const codeRef = createRef<HTMLTextAreaElement>();
    const router = useRouter()
    const lines = code.split('\n');
    const html = hljs.highlightAuto(code);

    useEffect(() => {
        if (codeRef.current)
            codeRef.current.innerHTML = html.value;
    }, [html, codeRef])

    useEffect(() => {
        const listener = (event: KeyboardEvent) => {
            if (event.code === "KeyN" && event.shiftKey === true) {
                event.preventDefault()
                router.push('/')
            }
        }

        document.addEventListener('keydown', listener)

        return () => {
            document.removeEventListener('keydown', listener)
        }
    }, [router])

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

export const getServerSideProps = async (context: GetServerSidePropsContext) => {
    const id = context.params?.id
    if (!id)
        return {
            redirect: {
                destination: '/',
                permanent: false,
            }
        }
    try {
        let data = await getData(id.toString())

        return {
            props: {
                code: data.data?.code,
            }
        }
    }
    catch {
        return {
            redirect: {
                destination: '/',
                permanent: false,
            }
        }
    }
}


export default Viewer
