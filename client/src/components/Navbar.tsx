import { Text, Container, useColorModeValue } from "@chakra-ui/react";


export default function Navbar() {
  return (
    <Container
      textAlign={"center"}
      backgroundColor={useColorModeValue("gray.400","")}
    >
      <Text
        color={useColorModeValue("yellow.600","yellow.300")}
        fontWeight={"bold"}
        fontSize={"4xl"}
      >
        SFPC YOUTH
      </Text>

      <Text
        color={"black"}
        fontWeight={"bold"}
        backgroundColor={"yellow.300"}
        fontSize={"4xl"}
      >
        Budget
      </Text>
    </Container>
  )
}