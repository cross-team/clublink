import React, { Component, ReactElement } from 'react';
import { UIFactory } from '../../UIFactory';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

import { Section } from '../../ui/Section';

interface IProps {
  isUserSignedIn?: boolean | undefined;
  uiFactory: UIFactory;
  handleOnShortLinkPageLoad: (offset: number, pageSize: number) => void;
  currentPagedShortLinks?: IPagedShortLinks | undefined;
}

export class FavoritesSection extends Component<IProps> {
  render(): ReactElement {
    return (
      <Section title={''}>
        {this.props.isUserSignedIn && (
          <div className={'user-short-links-section'}>
            {this.props.uiFactory.createUserShortLinksSection({
              onPageLoad: this.props.handleOnShortLinkPageLoad,
              pagedShortLinks: this.props.currentPagedShortLinks
            })}
          </div>
        )}
      </Section>
    );
  }
}
