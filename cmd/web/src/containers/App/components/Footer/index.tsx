import React from "react";
import { Box, Stack, Flex, useColorModeValue } from "@chakra-ui/core";
import { brandColors } from "src/theme/theme";
import { useLocale } from "src/hooks";
import { LOCALE } from "src/types";
import LanguageSwitcher from "src/components/common/LanguageSwitcher";
import ColorModeSwitcher from "src/components/common/ColorModeSwitcher";

const Footer = () => {
  const [locale, setLocale] = useLocale();

  const color = useColorModeValue(brandColors.light, brandColors.dark);

  const onLocaleChange = (locale: LOCALE) => setLocale(locale);

  return (
    <Flex
      as="footer"
      align="center"
      justify="space-between"
      wrap="wrap"
      padding="0.5rem"
      borderTopWidth="1px"
      borderColor={color.primary}
      mt={4}
    >
      <Box display="block" flexGrow={1}>
        <ColorModeSwitcher flex="1" />
      </Box>
      <Box display="block">
        <LanguageSwitcher
          flex="1"
          locale={locale}
          onLocaleChange={onLocaleChange}
        />
      </Box>
    </Flex>
  );
};

export default Footer;
