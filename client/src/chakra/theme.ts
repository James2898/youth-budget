import { extendTheme, ThemeConfig } from "@chakra-ui/react";
import {mode} from "@chakra-ui/theme-tools"

const config: ThemeConfig = {
    initialColorMode: "dark",
    useSystemColorMode: true
}

const theme = extendTheme({
    config,
    global: (props: any) => ({
        body: {
            backgroundColor: mode("gray.500","")
        }
    })
})

export default theme;